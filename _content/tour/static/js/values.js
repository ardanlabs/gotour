/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
'use strict';

angular.module('tour.values', []).

// List of modules with description and lessons in it.
value('tableOfContents', [{
    'id': 'welcome',
    'title': 'Welcome',
    'description': `<p>This is material for any intermediate-level developer who has some experience with other programming languages and wants to learn Go. We believe these classes are perfect for anyone who wants a jump start in learning Go or who wants a more thorough understanding of the language and its internals.</p>
    <p>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Language Mechanics',
    'description': `
        <p>This material covers all the language syntax, idioms, implementation and specification of the language. Once you are done with this material you will understand the mechanics of the language and mechanical sympathies the language has for both the hardware and operating system.</p>
        <p><i>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</i></p>
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
    'id': 'software-design',
    'title': 'Software Design',
    'description': `
        <p>This material covers the essential things you need to know about designing software in Go.</p>
        <p><i>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</i></p>
        `,
    'lessons': [
       'composition',
       'error-handling',
       'packaging',
    ]
},{
    'id': 'concurrency',
    'title': 'Concurrency',
    'description': `
        <p>This material covers all the concurrency aspects of the language. Once you are done with this material you will understand the concurrent mechanics of the language and mechanical sympathies the language has for both the hardware and operating system as it related to concurrency.</p>
        <p><i>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</i></p>
        `,
    'lessons': [
       'goroutines',
       'data_race',
       'context',
       'channels',
       'patterns'
    ]
},{
    'id': 'tooling',
    'title': 'Tooling',
    'description': `
        <p>This material covers a good portion of the tooling that comes with go. Specifically we cover testing and benchmarking. We also cover profiling memory and the scheduler. Finally we learn how to read stack traces.</p>
        <p><i>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</i></p>
        `,
    'lessons': [
       'testing',
       'benchmarks',
       'fuzzing',
       'profiling',
    ]
},{
    'id': 'packaging',
    'title': 'Packaging',
    'description': `
        <p>This material covers the essential things you need to know about the standard library and some important third party packages. Along the way you also learn about the most commonly used packages.</p>
        <p><i>Note: This material has been designed to be taught in a classroom environment. The code is well commented but missing some of the contextual concepts and ideas that will be covered in class.</i></p>
        `,
    'lessons': [
      'encoding',
      'io',
      'logging',
      'reflection',
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
