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
        'description': `<p>Bu, diğer programlama dilleri konusunda biraz deneyimi olan ve Go'yu öğrenmek isteyen orta düzey geliştiricilere yönelik bir materyaldir. Bu materyalin Go öğrenmeye hızlı bir başlangıç yapmak isteyen veya dili ve dilin iç yapısını daha kapsamlı anlamak isteyen herkes için mükemmel olduğuna inanıyoruz.</p>
    <p><b>Bir örneği ilk kez çalıştırdığınızda tamamlanması biraz zaman alabilir. Bunun nedeni tur hizmetinin çalışmıyor olmasıdır. Lütfen tamamlanması için biraz zaman verin.</b></p>
    `,
        'lessons': [
            'welcome',
        ]
    }, {
        'id': 'language-mechanics',
        'title': 'Dil Mekanikleri',
        'description': `
        <p>Bu materyal, dilin tüm sözdizimi, deyimleri, uygulaması ve dilin hem donanım hem de işletim sistemi için mekanik sempatisini kapsar. Bu materyali bitirdiğinizde, dilin mekaniklerini anlayacak ve dilin hem donanım hem de işletim sistemi için mekanik sempatisini kavrayacaksınız.</p>
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
        'title': 'Kompozisyon ve Arayüzler',
        'description': `
        <p>Bu materyal, kompozisyon ve arayüzler hakkında bilmeniz gereken pratik konuları kapsar.</p>
        `,
        'lessons': [
            'composition-grouping',
            'composition-assertions',
            'composition-pollution',
            'composition-mocking',
            'composition-decoupling',
            'error-handling',
        ]
    }, {
        'id': 'concurrency',
        'title': 'Paralellik',
        'description': `
        <p>Bu materyal, dilin tüm paralellik yönlerini kapsar. Bu materyali tamamladığınızda, dilin paralel mekaniğini anlayacak ve dilin donanım ve işletim sistemiyle ilgili olarak paralellikle ilgili mekanik duyarlılıklarını kavrayacaksınız.</p>
       `,
        'lessons': [
            'goroutines',
            'data_race',
            'context',
            'channels',
        ]
    }, {
        'id': 'generics',
        'title': 'Jenerikler',
        'description': `
        <p>Bu materyal, dilin tüm genel (generic) yönlerini kapsar. Generics, tip parametre listelerinin desteğiyle somut polimorfik fonksiyonları yazma yeteneği hakkında bir konsepttir.</p>
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
    }, {
        'id': 'algorithms',
        'title': 'Algoritmalar',
        'description': `
        <p>Bu materyal, ortak ve eğlenceli bir dizi algoritmayı uygulayan Go kodu sağlar.</p>
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
        'off': 'kapalı',
        'on': 'açık',
        'syntax': 'Sözdizimi-Vurgulama',
        'lineno': 'Satır-Numaraları',
        'reset': 'Slaytı Sıfırla',
        'format': 'Kaynak Kodu Biçimlendir',
        'kill': 'Programı Sonlandır',
        'run': 'Çalıştır',
        'compile': 'Derle ve Çalıştır',
        'more': 'Seçenekler',
        'toc': 'İçindekiler',
        'prev': 'Önceki',
        'next': 'Sonraki',
        'waiting': 'Uzak sunucu bekleniyor...',
        'errcomm': 'Uzak sunucu ile iletişim hatası.',
        'submit-feedback': 'Bu sayfa hakkında geri bildirim gönder',
        'search': 'İçerik ara',

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
                'Shift-Enter': function () {
                    $('#run').click();
                },
                'Ctrl-Enter': function () {
                    $('#format').click();
                },
                'PageDown': function () {
                    return false;
                },
                'PageUp': function () {
                    return false;
                },
            },
            // TODO: is there a better way to do this?
            // AngularJS values can't depend on factories.
            onChange: function () {
                if (window.codeChanged !== null) window.codeChanged();
            }
        }
    }).

    // mapping from the old paths (#42) to the new organization.
    // The values have been generated with the map.sh script in the tools directory.
    value('mapping', {
        '#1': '/variables/1', // Variables
    });
