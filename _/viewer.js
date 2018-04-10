angular.module("app", [])
    .controller('ViewerController', async function ($scope) {
        const res = await fetch("http://localhost:3000/result.json", {mode: 'no-cors'});
        const task = await res.json();
        document.t = task;
        console.log(task);
        // console.log(task.records);
        $scope.task = task;

        // headers
        $scope.headers = task.records.map(r => r.Name);

        // rows
        const rows = [];
        const len = task.misspells.length;
        for (let i = 0; i < len; i++) {
            const row = {
                misspell: task.misspells[i],
                correct: task.corrects[i],
                canCorrect: task.dictionary.includes(task.corrects[i]),
                candidatesOfMethods: task.records.map(r => {
                    const cs = r.Candidates[i];
                    return {
                        candidates: cs.join(", "),
                        len: cs.length,
                        hit: cs.includes(task.corrects[i])
                    }
                }),
                timesOfMethods: task.records.map(r => r.Times[i])
            };
            rows.push(row);
        }

        // stats
        $scope.times = task.records.map(r => r.Times.reduce((a, b) => a + b, 0));
        $scope.hitRates = [
            (() => {
                const total = rows.map(row => row.canCorrect).filter(v => v).length;
                return `${total}/${rows.length} - ${Math.round(total / rows.length * 10000) / 100}%`
            })()

        ].concat(task.records.map((r, i) => {
            const total = rows.map(row => row.candidatesOfMethods[i].hit).filter(b => b).length;
            return `${total}/${rows.length} - ${Math.round(total / rows.length * 10000) / 100}%`
        }));

        $scope.rows = rows;
        $scope.$apply();
    });
