var rootElement = document.documentElement;
rootElement.requestFullscreen = rootElement.requestFullscreen || rootElement.mozRequestFullScreen || rootElement.webkitRequestFullscreen || rootElement.msRequestFullscreen;
document.getElementById('button1').addEventListener('click', function() {
    rootElement.requestFullscreen();
});
var dum;

function INVISIBLE_IF_FULLSCREEN() {
    dum = setInterval(function() {
        if (document.fullscreen == true)
            document.getElementById('test_area').style.display = 'inline';
        else
            document.getElementById('test_area').style.display = 'none';
    }, 1);
}
window.document.onkeydown = function(evt) {
    if ((evt.which == 112) ||
        (evt.which == 71 && evt.ctrlKey == true) ||
        (evt.which == 70 && evt.ctrlKey == true) ||
        (evt.which == 114) ||
        (evt.which == 115) ||
        (evt.which == 82 && evt.ctrlKey == true) ||
        (evt.which == 116) ||
        (evt.which == 122) ||
        (evt.which == 73 && evt.ctrlKey == true && evt.shiftKey == true) ||
        (evt.which == 123) ||
        (evt.which == 9) ||
        (evt.which == 33) ||
        (evt.which == 34) ||
        (evt.which == 35) ||
        (evt.which == 36) ||
        (evt.which == 121 && evt.shiftKey == true) ||
        (evt.which == 13) ||
        (evt.which == 27) ||
        (evt.which == 32) ||
        (evt.which == 46)
    ) {
        evt.which = null;
        event.preventDefault();
    }
    document.oncontextmenu = function() { return false; }
}