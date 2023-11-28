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

// setLanguageSelectorChange initializes the language selector change event.
// It updates the URL and cookie based on the selected language.
function setLanguageSelectorChange() {
    const languageSelector = document.getElementById('languageSelector');
    languageSelector.addEventListener('change', (event) => {
        const currentUrl = window.location.pathname;
        const newLanguage = event.target.value;
        const newURL = replaceLanguageInUrl(currentUrl, newLanguage);

        window.location.href = newURL;
        document.cookie = `language-preference=${newLanguage};path=/;max-age=31536000;`;
    });
}

// setLanguageOptionBasedOnUrl sets the selected option of the language selector
// based on the language segment in the current URL path.
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

// replaceLanguageInUrl takes a URL and a new language as arguments,
// and returns a new URL with the language segment replaced by the new language.
function replaceLanguageInUrl(url, newLanguage) {
    return url.replace(/(\/tour\/)(eng|rus|per|grc|por|pol)(\/)/, `$1${newLanguage}$3`);
}

window.onload = function() {
    setLanguageOptionBasedOnUrl();
    setLanguageSelectorChange();
};