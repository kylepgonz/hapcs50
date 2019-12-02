window.onscroll = function() {myFunction()};
document.getElementById("navbar").style.background = "transparent";
var past = false;
function myFunction() {
  if (document.body.scrollTop > 500 || document.documentElement.scrollTop > 500) {
    past = true;
    document.getElementById("navbar").className = "navbar navbar-expand-lg fixed-top navbar-light slideUp";
    document.getElementById("navbar").style.background = "white";
  } else if ((document.body.scrollTop < 500 || document.documentElement.scrollTop < 500) && past == true) {
    document.getElementById("navbar").style.background = "transparent";
    document.getElementById("navbar").className = "navbar navbar-expand-lg fixed-top navbar-light slideDown";
    document.getElementById("navbar").style.color = "rgba(0,0,0,.9)";
    }
  }
