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
            redirectTo: '/tour/eng/list'
        }).
        when('/tour/eng/', {
            redirectTo: '/tour/eng/list'
        }).
        when('/tour/eng/list', {
            templateUrl: '/tour/eng/static/partials/list.html',
        }).
        when('/tour/eng/:lessonId/:pageNumber', {
            templateUrl: '/tour/eng/static/partials/editor.html',
            controller: 'EditorCtrl'
        }).
        when('/tour/eng/:lessonId', {
            redirectTo: '/tour/eng/:lessonId/1'
        }).
        otherwise({
            redirectTo: '/tour/eng/list'
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
        if (url.pathname != '/tour/eng/' || url.hash == '') {
            return;
        }
        $location.hash('');
        var m = mapping[url.hash];
        if (m === undefined) {
            console.log('unknown url, redirecting home');
            $location.path('/tour/eng/list');
            return;
        } 
      
        $location.path('/tour/eng' + m);
    });
});

window.scrollTo = (id) => {

    const element = document.getElementById(id);

    if(!element)
        return
    
    element.scrollIntoView();
}