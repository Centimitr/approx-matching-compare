angular.module("app", [])
    .controller('ViewerController', async function ($scope) {
        const res = await fetch("http://localhost:3000/result.json", {mode: 'no-cors'});
        const task = await res.json();
        document.t = task;
        console.log(task);
        console.log(task.records);
        task.records = task.records.filter(r => ['Neighbour(K=1)-1', 'nGram(N=2)-1', 'GED-1'].includes(r.Name));
        $scope.task = task;

        const nums = task.records.map(r => r.Candidates.map(css => css.length > 0 ? 0 : 1).reduce((a, b) => a + b, 0));


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
                    const cs = r.Candidates[i] || [];
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
        // $scope.times = task.records.map(r => r.Times.reduce((a, b) => a + b, 0));
        const recalls = $scope.recalls = [
            (() => {
                const total = rows.map(row => row.canCorrect).filter(v => v).length;
                return [`${total}/${rows.length} - ${Math.round(total / rows.length * 10000) / 100}%`, 0]
            })()

        ].concat(task.records.map((r, i) => {
            const hits = rows.map(row => row.candidatesOfMethods[i].hit).filter(b => b).length;
            const recall = hits / rows.length;
            return [`${hits}/${rows.length} - ${Math.round(recall * 10000) / 100}%`, recall]
        }));

        const precisions = $scope.precisions = task.records.map((r, i) => {
            const hits = rows.map(row => row.candidatesOfMethods[i].hit).filter(b => b).length;
            const predictions = rows.map(row => row.candidatesOfMethods[i].len).reduce((a, b) => a + b, 0);
            const precision = hits / predictions;
            return [`${hits}/${predictions} - ${Math.round(precision * 10000) / 100}%`, precision]
        });

        const f1s = $scope.f1s = task.records.map((r, i) => {
            const recall = recalls[i + 1][1];
            const precision = precisions[i][1];
            return 2 * (precision * recall) / (precision + recall);
        });
        $scope.rows = rows;
        $scope.notLoading = true;
        $scope.$apply();
    });
