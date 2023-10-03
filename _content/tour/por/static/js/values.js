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
    'description': `<p>Este material é para qualquer desenvolvedor de nível intermediário que tenha alguma experiência com outras linguagens de programação e quer aprender Go. Acreditamos que este material é perfeito para quem deseja começar a aprender Go ou deseja uma compreensão mais completa da linguagem e de seus componentes internos.</p>
    <p><b>A primeira vez que você executar um exemplo, pode levar algum tempo para ser concluído. Isso ocorre porque o serviço do tour pode não estar executando. Por favor, aguarde algum tempo para que ele conclua.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Mecânicas da Linguagem',
    'description': `
        <p>Este material cobre toda a sintaxe da linguagem, idiomas, implementação e especificação da linguagem. Depois de concluir este material, você entenderá a mecânica da linguagem e as simpatias mecânicas que a linguagem tem tanto para o hardware quanto para o sistema operacional.</p>
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
    'title': 'Composição e Interfaces',
    'description': `
        <p>Este material aborda coisas práticas que você precisa saber sobre composição e interfaces.</p>
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
    'title': 'Concorrência',
    'description': `
        <p>Este material aborda todos os aspectos de concorrência da linguagem. Ao concluir este material, você entenderá a mecânica de concorrência da linguagem e as simpatias mecânicas que a linguagem tem tanto para o hardware quanto para o sistema operacional no que se refere à concorrência.</p>
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
        <p>Este material aborda todos os aspectos sobre generics na linguagem. Generics é sobre fornecer a capacidade de escrever funções polimórficas concretas com o suporte de listas de parâmetros de tipo.</p>
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
    'title': 'Algoritmos',
    'description': `
        <p>Este material fornece códigos Go que implementam um conjunto de algoritmos comuns e divertidos.</p>
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
