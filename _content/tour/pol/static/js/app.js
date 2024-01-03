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
            redirectTo: '/tour/pol/list'
        }).
        when('/tour/pol/', {
            redirectTo: '/tour/pol/list'
        }).
        when('/tour/pol/list', {
            templateUrl: '/tour/pol/static/partials/list.html',
        }).
        when('/tour/pol/:lessonId/:pageNumber', {
            templateUrl: '/tour/pol/static/partials/editor.html',
            controller: 'EditorCtrl'
        }).
        when('/tour/pol/:lessonId', {
            redirectTo: '/tour/pol/:lessonId/1'
        }).
        otherwise({
            redirectTo: '/tour/pol/list'
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
        if (url.pathname != '/tour/pol/' || url.hash == '') {
            return;
        }
        $location.hash('');
        var m = mapping[url.hash];
        if (m === undefined) {
            console.log('unknown url, redirecting home');
            $location.path('/tour/pol/list');
            return;
        } 
      
        $location.path('/tour/pol' + m);
    });
});

window.scrollTo = (id) => {

    const element = document.getElementById(id);

    if(!element)
        return
    
    element.scrollIntoView();
}