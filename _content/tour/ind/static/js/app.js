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
            redirectTo: '/tour/ind/list'
        }).
        when('/tour/ind/', {
            redirectTo: '/tour/ind/list'
        }).
        when('/tour/ind/list', {
            templateUrl: '/tour/ind/static/partials/list.html',
        }).
        when('/tour/ind/:lessonId/:pageNumber', {
            templateUrl: '/tour/ind/static/partials/editor.html',
            controller: 'EditorCtrl'
        }).
        when('/tour/ind/:lessonId', {
            redirectTo: '/tour/ind/:lessonId/1'
        }).
        otherwise({
            redirectTo: '/tour/ind/list'
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
        if (url.pathname != '/tour/ind/' || url.hash == '') {
            return;
        }
        $location.hash('');
        var m = mapping[url.hash];
        if (m === undefined) {
            console.log('unknown url, redirecting home');
            $location.path('/tour/ind/list');
            return;
        } 
      
        $location.path('/tour/ind' + m);
    });
});

window.scrollTo = (id) => {

    const element = document.getElementById(id);

    if(!element)
        return
    
    element.scrollIntoView();
}