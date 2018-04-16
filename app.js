var app = window.angular.module('Fib', [])

app.factory('numberFactory', [function(){
                         var o = {
                         sequence: []
                         };
                         return o;
                         }])

app.controller('mainCtrl', [
    '$scope','$http', 'numberFactory',
    function($scope, $http, numberFactory) {
        var input = [];
        $scope.getFib = function() {
            input = [];
            if($scope.formContent === '') {return;}
            console.log('localhost:8080/fibonacci/' + $scope.formContent)
            $http.get('http://localhost:8080/fibonacci/' + $scope.formContent).success(function(data){
                    angular.copy(data.Sequence, input);
            });
            $scope.sequence = input;
        };
    }
]);



