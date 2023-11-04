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
    'description': `<p>Jest to materiał dla każdego programisty na średnim poziomie, który ma już pewne doświadczenie z innymi językami programowania i chce nauczyć się Go. Wierzymy, że ten materiał jest idealny dla każdego, kto chce szybko rozpocząć naukę Go lub chce dokładniej zrozumieć język i jego elementy wewnętrzne.</p>
    <p><b>Za pierwszym razem, gdy uruchamiasz dany przykład, może minąć trochę czasu zanim zostanie on zakończony. Wynika to z tego, że usługa wycieczkowa po Go, może nie być aktualnie uruchomiona. Daj jej trochę czasu na zakończenie.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Mechanika języka',
    'description': `
        <p>Ten materiał obejmuje składnię języka, idiomatykę, implementację i specyfikację języka. Po ukończeniu tego materiału zrozumiesz mechanikę języka oraz jego mechaniczne zależności wobec sprzętu i systemu operacyjnego.</p>
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
    'title': 'Kompozycja i interfejsy',
    'description': `
        <p>Ten materiał obejmuje praktyczne informacje, które musisz wiedzieć na temat kompozycji i interfejsów.</p>
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
    'title': 'Współbieżność',
    'description': `
        <p>Ten materiał obejmuje wszystkie aspekty współbieżności języka. Po ukończeniu tego materiału zrozumiesz mechanikę współbieżności języka oraz jego mechaniczne zależności wobec sprzętu i systemu operacyjnego w kontekście współbieżności.</p>
       `,
    'lessons': [
        'goroutines',
        'data_race',
        'context',
        'channels',
    ]
},{
    'id': 'generics',
    'title': 'Generyki',
    'description': `
        <p>Ten materiał obejmuje wszystkie aspekty ogólnych (generics) języka. Ogólne (generics) to możliwość tworzenia konkretnych funkcji polimorficznych z wykorzystaniem listy parametrów typu.</p>
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
    'title': 'Algorytmy',
    'description': `
        <p>Ten materiał dostarcza kod Go, implementujący popularny i zabawny zestaw algorytmów.</p>
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
