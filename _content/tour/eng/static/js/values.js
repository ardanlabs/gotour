/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'welcome',
    'title': '',
    'description': `<p>This is material for any intermediate-level developer who has some experience with other programming languages and wants to learn Go. We believe this material is perfect for anyone who wants a jump start in learning Go or who wants a more thorough understanding of the language and its internals.</p>
    <p><b>The first time you run an example it may take some time to complete. This is because the tour service might not be running. Please give it some time to complete.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Language Mechanics',
    'description': `
        <p>This material covers all the language syntax, idioms, implementation and specification of the language. Once you are done with this material you will understand the mechanics of the language and mechanical sympathies the language has for both the hardware and operating system.</p>
        `,
    'lessons': [
        'variables',
        'struct-types', 
        'pointers', 
        'constants',
        'functions',
        'arrays',
        'slices',
        'maps',
        'methods',
        'interfaces',
        'embedding',
        'exporting',
    ]
}, {
    'id': 'composition-interfaces',
    'title': 'Composition and Interfaces',
    'description': `
        <p>This material covers the practical things you need to know about composition and interfaces.</p>
        `,
    'lessons': [
        'composition-grouping',
        'composition-assertions',
        'composition-pollution',
        'composition-mocking',
        'composition-decoupling',
        'error-handling',
    ]
},{
    'id': 'concurrency',
    'title': 'Concurrency',
    'description': `
        <p>This material covers all the concurrency aspects of the language. Once you are done with this material you will understand the concurrent mechanics of the language and mechanical sympathies the language has for both the hardware and operating system as it related to concurrency.</p>
       `,
    'lessons': [
        'goroutines',
        'data_race',
        'context',
        'channels',
    ]
},{
    'id': 'generics',
    'title': 'Generics',
    'description': `
        <p>This material covers all the generics aspects of the language. Generics is about providing the ability to write concrete polymorphic functions with the support of type parameter lists.</p>
       `,
    'lessons': [
        'generics-basics',
        'generics-underlying-types',
        'generics-struct-types',
        'generics-behavior-constraints',
        'generics-type-constraints',
        'generics-multi-type-params',
        'generics-slice-constraints',
        'generics-channels',
        'generics-hash-table',
    ]
},{
    'id': 'algorithms',
    'title': 'Algorithms',
    'description': `
        <p>This material provides Go code implementing a common and fun set of algorithms.</p>
       `,
    'lessons': [
        'algorithms-bits-seven',
        'algorithms-strings',
        'algorithms-numbers',
        'algorithms-slices',
        'algorithms-sorting',
        'algorithms-data',
        'algorithms-searches',
        'algorithms-fun',
    ]
}]).
// translation
value('translation', {
    'off': 'off',
    'on': 'on',
    'syntax': 'Syntax-Highlighting',
    'lineno': 'Line-Numbers',
    'reset': 'Reset Slide',
    'format': 'Format Source Code',
    'kill': 'Kill Program',
    'run': 'Run',
    'compile': 'Compile and Run',
    'more': 'Options',
    'toc': 'Table of Contents',
    'prev': 'Previous',
    'next': 'Next',
    'waiting': 'Waiting for remote server...',
    'errcomm': 'Error communicating with remote server.',
    'submit-feedback': 'Send feedback about this page',
    'search': 'Search for content',

    // GitHub issue template: update repo and messaging when translating.
    'github-repo': 'github.com/ardanlabs/gotour',
    'issue-title': 'tour: [REPLACE WITH SHORT DESCRIPTION]',
    'issue-message': 'Change the title above to describe your issue and add your feedback here, including code if necessary',
    'context': 'Context',
}).

// Config for codemirror plugin
value('ui.config', {
    codemirror: {
        mode: 'text/x-go',
        matchBrackets: true,
        lineNumbers: true,
        autofocus: true,
        indentWithTabs: true,
        indentUnit: 4,
        tabSize: 4,
        lineWrapping: true,
        extraKeys: {
            'Shift-Enter': function() {
                $('#run').click();
            },
            'Ctrl-Enter': function() {
                $('#format').click();
            },
            'PageDown': function() {
                return false;
            },
            'PageUp': function() {
                return false;
            },
        },
        // TODO: is there a better way to do this?
        // AngularJS values can't depend on factories.
        onChange: function() {
            if (window.codeChanged !== null) window.codeChanged();
        }
    }
}).

// mapping from the old paths (#42) to the new organization.
// The values have been generated with the map.sh script in the tools directory.
value('mapping', {
    '#1': '/variables/1', // Variables
});
