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
    'description': `
    <p>Αυτό είναι υλικό για κάθε developer ενδιάμεσου επιπέδου με κάποια εμπειρία σε άλλες γλώσσες προγραμματισμού που θέλει να μάθει την Go. Πιστεύουμε ότι αυτό το υλικό είναι τέλειο για οποιονδήποτε θέλει να αποκτήσει ένα προβάδισμα στην κατανόηση της Go ή που θέλει μια βαθύτερη κατανόηση της γλώσσας και των εσωτερικών της μηχανισμών.</p>
    <p><b>Την πρώτη φορά που θα τρέξετε ένα παράδειγμα μπορεί να πάρει κάποιο χρόνο για να ολοκληρωθεί. Αυτό συμβαίνει γιατί η υπηρεσία της επισκόπησης ενδέχεται να μην λειτουργεί ήδη. Παρακαλούμε, δώστε της λίγο χρόνο προκειμένου να ολοκληρωθεί.</b></p>
    `,
    'lessons': [
        'welcome',
    ]
}, {
    'id': 'language-mechanics',
    'title': 'Μηχανισμοί της Γλώσσας',
    'description': `<p>Αυτό το υλικό καλύπτει ολόκληρο το συντακτικό της γλώσσας, τα ιδιώματα, την υλοποίηση και τις προδιαγραφές της γλώσσας. Όταν ολοκληρώσετε το υλικό θα κατανοείτε τους μηχανισμούς της γλώσσας και τις μηχανικές συμπάθειες που έχει η γλώσσα τόσο για το υλικό όσο και για το λειτουργικό σύστημα.</p>
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
    'title': 'Σύνθεση και Διεπαφές',
    'description': `
        <p>Αυτό το υλικό καλύπτει τα πρακτικά πράγματα που χρειάζεται να ξέρετε σχετικά με την σύνθεση και τις διεπαφές.</p>
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
    'title': 'Ταυτόχρονη Εκτέλεση',
    'description': `
        <p>Αυτό το υλικό καλύπτει όλες τις πτυχές της γλώσσας που αφορούν την ταυτόχρονη εκτέλεση. Όταν ολοκληρώσετε αυτό το υλικό θα κατανοείτε τους μηχανισμούς της παράλληλης εκτέλεσης της γλώσσας και τις μηχανικές συμπάθειες που έχει η γλώσσα τόσο για το υλικό όσο και για το λειτουργικό σύστημα στα ζητήματα που αφορούν την ταυτόχρονη εκτέλεση.</p>
       `,
    'lessons': [
        'goroutines',
        'data_race',
        'context',
        'channels',
    ]
},{
    'id': 'generics',
    'title': 'Γενικός Προγραμματισμός',
    'description': `
        <p>Αυτό το υλικό καλύπτει όλες της πτυχές της γλώσσας που σχετίζονται με τον γενικό προγραμματισμό. Ο γενικός προγραμματισμός αφορούν την δυνατότητα συγγραφής πραγματικών πολύμορφων συναρτήσεων με την υποστήριξη παραμέτρων τύπων.</p>
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
    'title': 'Αλγόριθμοι',
    'description': `
        <p>Αυτό το υλικό παρέχει κώδικα της Go που υλοποιεί ένα σύνολο κοινών και διασκεδαστικών αλγορίθμων.</p>
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
    'syntax': 'Υπογράμμιση Σύνταξης',
    'lineno': 'Αριθμοί Γραμμής',
    'reset': 'Επαναφορά Slide',
    'format': 'Μορφοποίηση Πηγαίου Κώδικα',
    'kill': 'Διακοπή Προγράμματος',
    'run': 'Εκτέλεση',
    'compile': 'Μεταγλώττιση και Εκτέλεση',
    'more': 'Επιλογές',
    'toc': 'Πίνακας Περιεχομένων',
    'prev': 'Προηγούμενο',
    'next': 'Επόμενο',
    'waiting': 'Αναμένοντας τον απομακρυσμένο server...',
    'errcomm': 'Σφάλμα επικοινωνίας με τον απομακρυσμένο server.',
    'submit-feedback': 'Υποβολή Σχολίων σχετικά με την σελίδα',
    'search': 'Αναζήτηση για περιεχόμενο',

    // GitHub issue template: update repo and messaging when translating.
    'github-repo': 'github.com/ardanlabs/gotour',
    'issue-title': 'επισκόπηση: [ΑΝΤΙΚΑΤΑΣΤΕIΣΤΕ ΜΕ ΣΥΝΤΟΜΗ ΠΕΡΙΓΡΑΦΗ]',
    'issue-message': 'Αλλάξτε τον παραπάνω τίτλο για να περιγράψετε το πρόβλημα που αντιμετωπίζετε και προσθέστε το σχόλιο σας εδώ, μαζί με τον όποιο κώδικα αν αυτό είναι αναγκαίο',
    'context': 'Πλαίσιο',
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
