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
      description: `<p>Dies ist Material für jeden Entwickler mittlerer Stufe, der Erfahrung mit anderen Programmiersprachen hat und Go lernen möchte. Wir glauben, dass dieses Material perfekt für alle ist, die einen schnellen Einstieg in das Erlernen von Go wünschen oder ein tieferes Verständnis der Sprache und ihrer Interna wünschen.</p>
    <p><b>Wenn Sie ein Beispiel zum ersten Mal ausführen, kann es einige Zeit dauern, bis es fertig ist. Dies liegt daran, dass der Tourservice möglicherweise nicht ausgeführt wird. Bitte geben Sie etwas Zeit für die Fertigstellung.</b></p>
    `,
      lessons: ["welcome"],
    },
    {
      id: "language-mechanics",
      title: "Language Mechanics",
      description: `
        <p>Dieses Material behandelt die gesamte Sprachsyntax, Idiome, Implementierung und Spezifikation der Sprache. Wenn Sie mit diesem Material fertig sind, werden Sie die Mechanik der Sprache und die mechanischen sympathies verstehen, die die Sprache sowohl für die Hardware als auch für das Betriebssystem hat.</p>
        `,
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
      title: "Composition and Interfaces",
      description: `
        <p>Dieses Material behandelt die praktischen Dinge, die Sie über Komposition und Schnittstellen wissen müssen.</p>
        `,
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
      title: "Concurrency",
      description: `
        <p>Dieses Material deckt alle Gleichzeitigkeitsaspekte der Sprache ab. Wenn Sie mit diesem Material fertig sind, werden Sie die Gleichzeitigkeitsmechanik der Sprache und die mechanischen Sympathien verstehen, die die Sprache sowohl für die Hardware als auch für das Betriebssystem hat, wenn es um Gleichzeitigkeit geht.</p>
       `,
      lessons: ["goroutines", "data_race", "context", "channels"],
    },
    {
      id: "generics",
      title: "Generics",
      description: `
        <p>Dieses Material behandelt alle generischen Aspekte der Sprache. Bei den Generika geht es um die Möglichkeit, konkrete polymorphe Funktionen mit Unterstützung von Typparameterlisten zu schreiben.</p>
       `,
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
      title: "Algorithms",
      description: `
        <p>Dieses Material bietet Go-Code, der eine Reihe von gängigen und unterhaltsamen Algorithmen implementiert.</p>
       `,
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
