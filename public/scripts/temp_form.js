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




(function( $, window, undefined ) {


    var bs = {
        close: $(".icon-remove-sign"),
        searchform: $("form"),
        canvas: $("body"),
        dothis: $('.dosearch')
    };

    bs.dothis.on('click', function() {
        $(this).toggleClass('hidden');
        bs.searchform.toggleClass('active');
        bs.searchform.find('input').focus();
        //bs.searchform.toggleClass('hid');
        //bs.canvas.toggleClass('overlay');
    });

    bs.close.on('click', function() {
        bs.searchform.toggleClass('active');
        //bs.canvas.removeClass('overlay');
    });

})( jQuery, window );