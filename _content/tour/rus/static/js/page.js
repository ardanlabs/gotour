window.transport = {{.Transport}}();
window.socketAddr = "{{.SocketAddr}}";

// highlight applies a highlight effect to a specific element determined by
// the selector parameter.
function highlight(selector) {
    var speed = 50;
    var obj = $(selector).stop(true, true)
    for (var i = 0; i < 5; i++) {
        obj.addClass("highlight", speed)
        obj.delay(speed)
        obj.removeClass("highlight", speed)
    }
}

// highlightAndClick highlights an element (as per the highlight function)
// and then simulates a click on the element after a delay.
function highlightAndClick(selector) {
    highlight(selector);
    setTimeout(function() {
        $(selector)[0].click()
    }, 750);
}

// click simulates a click event on the element determined by the selector
// parameter.
function click(selector) {
    $(selector)[0].click();
}

// setThemeAttribute sets the theme attribute on the document element and
// updates the theme preference cookie.
// Theme to set ('auto', 'dark', or 'light').
function setThemeAttribute(theme) {
    document.documentElement.setAttribute('data-theme', theme);
    let domain = '';
    if (location.hostname === 'tour.ardanlabs.com') {
        // Include subdomains to apply the setting to pkg.go.dev.
        domain = 'domain=.ardanlabs.com;';
    }

    document.cookie = `prefers-color-scheme=${theme};${domain}path=/;max-age=31536000;`;
}

// setInitialTheme sets the initial theme based on the OS preference when
// the page is loaded.
// If the OS preference cannot be determined, defaults to 'auto'.
function setInitialTheme() {
    let initialTheme = 'auto';

    if (window.matchMedia) {
        const prefersDarkScheme = window.matchMedia("(prefers-color-scheme: dark)").matches;
        const prefersLightScheme = window.matchMedia("(prefers-color-scheme: light)").matches;

        initialTheme = prefersDarkScheme ? 'dark' : prefersLightScheme ? 'light' : 'auto';
    }

    setThemeAttribute(initialTheme);
}

// Set the initial theme based on OS preference when the page loads.
setInitialTheme();

// toggleTheme toggles the theme attribute between
// 'auto', 'dark', and 'light' based on the current setting.
function toggleTheme() {
    const currentTheme = document.documentElement.getAttribute('data-theme');
    let nextTheme = currentTheme === 'dark' ? 'light' : currentTheme === 'light' ? 'auto' : 'dark';

    setThemeAttribute(nextTheme);
}

// setThemeButtons adds click event listeners to all elements with the class
// 'js-toggleTheme' to toggle the theme when clicked.
function setThemeButtons() {
    for (const el of document.querySelectorAll('.js-toggleTheme')) {
        el.addEventListener('click', () => {
            toggleTheme();
        });
    }
}

setThemeButtons();

function setLanguageSelectorChange() {
    const languageSelector = document.getElementById('languageSelector');
    languageSelector.addEventListener('change', (event) => {
        window.location.href = `/tour/${event.target.value}/`;
        document.cookie = "language-preference=" + event.target.value + ";path=/;max-age=31536000;";
    });
}

function setLanguageOptionBasedOnUrl() {
    const languageSelector = document.getElementById('languageSelector');
    const currentUrl = window.location.pathname;

    for (let option of languageSelector.options) {
        if (currentUrl.includes(option.value)) {
            option.selected = true;
            break;
        }
    }
}

window.onload = function() {
    setLanguageOptionBasedOnUrl();
    setLanguageSelectorChange();
};