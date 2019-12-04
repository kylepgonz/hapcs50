window.onscroll = function() {myFunction()};
document.getElementById("navbar").style.background = "transparent";
var past = false;
function myFunction() {
  if (document.body.scrollTop > 500 || document.documentElement.scrollTop > 500) {
    past = true;
    document.getElementById("navbar").className = "navbar navbar-expand-md fixed-top navbar-light slideUp";

    document.getElementById("navbar").style.background = "white";
  } else if ((document.body.scrollTop < 500 || document.documentElement.scrollTop < 500) && past == true) {
    document.getElementById("navbar").style.background = "transparent";
<<<<<<< HEAD
    document.getElementById("navbar").className = "navbar navbar-expand-md fixed-top navbar-light slideDown";
=======
    document.getElementById("navbar").className = "navbar navbar-expand-lg fixed-top navbar-light slideDown";
>>>>>>> a4a7e94834a5f780c54ba03d78ecc8a2b58ddb7c
    document.getElementById("navbar").style.color = "rgba(0,0,0,.9)";
    }
  }
