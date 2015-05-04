var trigger = document.getElementById("hamburger");
var menu = document.getElementById("sidebar");
var menuIsClosed = false;
trigger.onclick = burgerTime;

function burgerTime() {
    if (menuIsClosed == true) {
        trigger.className = 'hamburglar menu-open-button';
        menu.className = 'sidebar';
        menuIsClosed = false;
    } else {
        trigger.className = 'hamburglar menu-close-button';
        menu.className = 'sidebar sidebar-open';
        menuIsClosed = true;
    }
}