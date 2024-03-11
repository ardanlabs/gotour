/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour', ['ui', 'tour.services', 'tour.controllers', 'tour.directives', 'tour.values', 'ng']).

config(['$routeProvider', '$locationProvider',
    function($routeProvider, $locationProvider) {
        $routeProvider.
        when('/tour/', {
            redirectTo: '/tour/fre/list'
        }).
        when('/tour/fre/', {
            redirectTo: '/tour/fre/list'
        }).
        when('/tour/fre/list', {
            templateUrl: '/tour/fre/static/partials/list.html',
        }).
        when('/tour/fre/:lessonId/:pageNumber', {
            templateUrl: '/tour/fre/static/partials/editor.html',
            controller: 'EditorCtrl'
        }).
        when('/tour/fre/:lessonId', {
            redirectTo: '/tour/fre/:lessonId/1'
        }).
        otherwise({
            redirectTo: '/tour/fre/list'
        });

        $locationProvider.html5Mode(true).hashPrefix('!');
    }
]).

// handle mapping from old paths (#42) to the new organization.
run(function($rootScope, $location, mapping) {
    $rootScope.$on( "$locationChangeStart", function(event, next) {
        var url = document.createElement('a');
        url.href = next;
        if (url.pathname != '/tour/' || url.hash == '') {
            return;
        }
        if (url.pathname != '/tour/fre/' || url.hash == '') {
            return;
        }
        $location.hash('');
        var m = mapping[url.hash];
        if (m === undefined) {
            console.log('unknown url, redirecting home');
            $location.path('/tour/fre/list');
            return;
        } 
      
        $location.path('/tour/fre' + m);
    });
});

window.scrollTo = (id) => {

    const element = document.getElementById(id);

    if(!element)
        return
    
    element.scrollIntoView();
}