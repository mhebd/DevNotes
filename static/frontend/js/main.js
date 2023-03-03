// all js function loader code

// nivo slider loader start
$(window).load(function () {
    $('#slider').nivoSlider();
  });

  var _gaq = _gaq || [];
  _gaq.push(['_setAccount', 'UA-36251023-1']);
  _gaq.push(['_setDomainName', 'jqueryscript.net']);
  _gaq.push(['_trackPageview']);

  (function () {
    var ga = document.createElement('script');
    ga.type = 'text/javascript';
    ga.async = true;
    ga.src = ('https:' == document.location.protocol ? 'https://ssl' : 'http://www') +
    '.google-analytics.com/ga.js';
    var s = document.getElementsByTagName('script')[0];
    s.parentNode.insertBefore(ga, s);
  })();
// nivo slider end

// navbar sub menu show code
$('ul li').click(function(){
  $('.first-d').toggleClass('dropdown-show')
});
$('ul li ul li').click(function(){
  $('.second-d').toggleClass('sub-dropdown-show')
});

// navbar scroll efect
function Scroll() {
  var nav = document.getElementById('top-navigation');
  var top = document.getElementById('go-top');
  var ypos = window.pageYOffset;
  if (ypos > 300) {
        nav.classList.add('scroll');
        top.classList.add('scroll');
  } else {
      nav.classList.remove('scroll');
      top.classList.remove('scroll');
  }
}
window.addEventListener("scroll", Scroll);

// cool scrolling effect
$(function() {
  var topoffset = 70; //variable for menu height

  //Use smooth scrolling when clicking on navigation
  $('.go-top a').click(function() {
  if (location.pathname.replace(/^\//,'') === 
    this.pathname.replace(/^\//,'') && 
    location.hostname === this.hostname) {
    var target = $(this.hash);
    target = target.length ? target : $('[name=' + this.hash.slice(1) +']');
    if (target.length) {
    $('html,body').animate({
      scrollTop: target.offset().top-topoffset
    }, 500);
    return false;
    } //target.length
  } //click function
  }); //smooth scrolling
});

// wow activate code

new WOW ().init();
