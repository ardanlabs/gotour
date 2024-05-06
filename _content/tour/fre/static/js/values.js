/* Copyright 2012 The Go Authors.   All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
"use strict";

angular
  .module("tour.values", [])
  // List of modules with description and lessons in it.
  .value("tableOfContents", [
    {
      id: "welcome",
      title: "",
      description: `<p>Ce matériel s'adresse à tout développeur de niveau intermédiaire ayant une certaine expérience d'autres langages de programmation et souhaitant apprendre Go. Nous pensons que ce matériel est parfait pour tous ceux qui veulent commencer à apprendre Go ou qui veulent une compréhension plus approfondie du langage et de ses aspects internes.</p>
      <p><b>La première fois que vous exécutez un exemple, l'exécution peut prendre un certain temps. Cela s'explique par le fait que le service du tour de Go n'est peut-être pas en cours d'exécution. Veuillez lui laisser un peu de temps pour s'exécuter.</b></p>`,
      lessons: ["welcome"],
    },
    {
      id: "language-mechanics",
      title: "Mécaniques du Langage",
      description: `<p>Ce matériel couvre toute la syntaxe du langage, les idiomes, la mise en œuvre et la spécification du langage. Une fois que vous aurez terminé ce matériel, vous comprendrez les mécanismes du langage et les sympathies mécaniques que le langage a pour le matériel et le système d'exploitation.</p>`,
      lessons: [
        "variables",
        "struct-types",
        "pointers",
        "constants",
        "functions",
        "arrays",
        "slices",
        "maps",
        "methods",
        "interfaces",
        "embedding",
        "exporting",
      ],
    },
    {
      id: "composition-interfaces",
      title: "Composition et Interfaces",
      description: `<p>Ce matériel couvre les aspects pratiques que vous devez connaître sur la composition et les interfaces.</p>`,
      lessons: [
        "composition-grouping",
        "composition-assertions",
        "composition-pollution",
        "composition-mocking",
        "composition-decoupling",
        "error-handling",
      ],
    },
    {
      id: "concurrency",
      title: "Concurrence",
      description: `<p>Ce matériel couvre tous les aspects de la concurrence du langage. Une fois que vous aurez terminé ce matériel, vous comprendrez la mécanique de concurrence du langage et les sympathies mécaniques que le langage a pour le matériel et le système d'exploitation en ce qui concerne la concurrence.</p>`,
      lessons: ["goroutines", "data_race", "context", "channels"],
    },
    {
      id: "generics",
      title: "Generics",
      description: `<p>Ce matériel couvre tous les aspects des génériques du langage. Les génériques permettent d'écrire des fonctions polymorphes concrètes à l'aide de listes de paramètres de type.</p>`,
      lessons: [
        "generics-basics",
        "generics-underlying-types",
        "generics-struct-types",
        "generics-behavior-constraints",
        "generics-type-constraints",
        "generics-multi-type-params",
        "generics-slice-constraints",
        "generics-channels",
        "generics-hash-table",
      ],
    },
    {
      id: "algorithms",
      title: "Algorithmes",
      description: `<p>Ce matériel fournit du code Go mettant en œuvre un ensemble d'algorithmes communs et amusants.</p>`,
      lessons: [
        "algorithms-bits-seven",
        "algorithms-strings",
        "algorithms-numbers",
        "algorithms-slices",
        "algorithms-sorting",
        "algorithms-data",
        "algorithms-searches",
        "algorithms-fun",
      ],
    },
  ])
  // translation
  .value("translation", {
    off: "off",
    on: "on",
    syntax: "Syntax-Highlighting",
    lineno: "Line-Numbers",
    reset: "Reset Slide",
    format: "Format Source Code",
    kill: "Kill Program",
    run: "Run",
    compile: "Compile and Run",
    more: "Options",
    toc: "Table of Contents",
    prev: "Previous",
    next: "Next",
    waiting: "Waiting for remote server...",
    errcomm: "Error communicating with remote server.",
    "submit-feedback": "Send feedback about this page",
    search: "Search for content",

    // GitHub issue template: update repo and messaging when translating.
    "github-repo": "github.com/ardanlabs/gotour",
    "issue-title": "tour: [REPLACE WITH SHORT DESCRIPTION]",
    "issue-message":
      "Change the title above to describe your issue and add your feedback here, including code if necessary",
    context: "Context",
  })
  // Config for codemirror plugin
  .value("ui.config", {
    codemirror: {
      mode: "text/x-go",
      matchBrackets: true,
      lineNumbers: true,
      autofocus: true,
      indentWithTabs: true,
      indentUnit: 4,
      tabSize: 4,
      lineWrapping: true,
      extraKeys: {
        "Shift-Enter": function () {
          $("#run").click();
        },
        "Ctrl-Enter": function () {
          $("#format").click();
        },
        PageDown: function () {
          return false;
        },
        PageUp: function () {
          return false;
        },
      },
      // TODO: is there a better way to do this?
      // AngularJS values can't depend on factories.
      onChange: function () {
        if (window.codeChanged !== null) window.codeChanged();
      },
    },
  })
  // mapping from the old paths (#42) to the new organization.
  // The values have been generated with the map.sh script in the tools directory.
  .value("mapping", {
    "#1": "/variables/1", // Variables
  });
