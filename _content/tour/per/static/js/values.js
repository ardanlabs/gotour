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
    'description': `<p>این مطلب برای هر توسعه دهنده سطح متوسطی است که تجربه ای با سایر زبان های برنامه نویسی دارد و می خواهد Go را یاد بگیرد. ما معتقدیم که این مطالب برای هرکسی که می‌خواهد شروعی سریع در یادگیری Go داشته باشد یا می‌خواهد درک کامل‌تری از زبان و درونیات آن داشته باشد، عالی است.</p>
    <p><b>اولین باری که یک مثال را اجرا می کنید ممکن است کمی طول بکشد تا کامل شود. این به این دلیل است که ممکن است سرویس تور در حال اجرا نباشد. لطفا کمی زمان بدهید تا کامل شود.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'مکانزیم زبان',
    'description': `
        <p>این متریال تمام نکات نحوه نوشتن، اصطلاحات، پیاده‌سازی، و مشخصات زبان را پوشش می‌دهد. بعد از اتمام این مواد، شما مکانیک‌های زبان و همدلی‌های مکانیکی آن با سخت‌افزار و سیستم عامل را خواهید فهمید.</p>
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
    'title': 'Composition اینترفیس و',
    'description': `
        <p>این مطلب شامل مسائل عملیاتی است که باید درباره ترکیب اینترفیس ها و Composition مختلف آنها بدانید.</p>
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
    'title': 'همزمانی',
    'description': `
        <p>این مطلب تمام جنبه های همروندی زبان را پوشش می دهد. پس از خواندن این مطلب، شما مکانیزم های همروندی زبان و همچنین همدردی های مکانیکی که زبان در ارتباط با سخت افزار و سیستم عامل در زمینه همروندی دارد را درک خواهید کرد.</p>
       `,
    'lessons': [
        'goroutines',
        'data_race',
        'context',
        'channels',
    ]
},{
    'id': 'generics',
    'title': 'جنریک ها',
    'description': `
        <p>این مطلب تمام جنبه های جنریک زبان را پوشش می دهد. جنریک ها در مورد ارائه قابلیت نوشتن توابع پلی مورفیک محکم با پشتیبانی از لیست پارامترهای نوع است.</p>
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
    'title': 'الگوریتم ها',
    'description': `
        <p>این مطلب کد Go را ارائه می دهد که یک مجموعه معمول و جذاب از الگوریتم ها را پیاده سازی می کند.</p>
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
