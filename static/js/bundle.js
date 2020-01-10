(function(){function r(e,n,t){function o(i,f){if(!n[i]){if(!e[i]){var c="function"==typeof require&&require;if(!f&&c)return c(i,!0);if(u)return u(i,!0);var a=new Error("Cannot find module '"+i+"'");throw a.code="MODULE_NOT_FOUND",a}var p=n[i]={exports:{}};e[i][0].call(p.exports,function(r){var n=e[i][1][r];return o(n||r)},p,p.exports,r,e,n,t)}return n[i].exports}for(var u="function"==typeof require&&require,i=0;i<t.length;i++)o(t[i]);return o}return r})()({1:[function(require,module,exports){
class App extends React.Component {
  render() {
    return React.createElement("div", null, React.createElement(Home, null));
  }

}

class Home extends React.Component {
  render() {
    return React.createElement("div", {
      className: "container"
    }, React.createElement("div", {
      className: "row"
    }, React.createElement("div", {
      className: "col-10 text-center pt-5 pl-5"
    }, React.createElement("div", null, React.createElement("img", {
      className: "headshot",
      src: "/static/img/headshot.jpg",
      alt: "cyrus javan headshot"
    })), React.createElement("h1", {
      className: "my-3"
    }, "Cyrus Javan"), React.createElement("h2", {
      className: "my-3"
    }, "Software Engineer"), React.createElement("h2", {
      className: "my-3"
    }, "Currently helping build the world's largest online counseling platform at ", React.createElement("a", {
      href: "https://www.betterhelp.com"
    }, "BetterHelp"))), React.createElement("div", {
      className: "col-2 text-center py-5"
    }, React.createElement("p", null, React.createElement("a", {
      className: "social-icon",
      href: "https://github.com/CyrusJavan"
    }, React.createElement("i", {
      className: "fa fa-github"
    }))), React.createElement("p", null, React.createElement("a", {
      className: "social-icon",
      href: "mailto:javan.cyrus@gmail.com"
    }, React.createElement("i", {
      className: "fa fa-envelope"
    }))), React.createElement("p", null, React.createElement("a", {
      className: "social-icon",
      href: "https://www.linkedin.com/in/cyrusjavan"
    }, React.createElement("i", {
      className: "fa fa-linkedin"
    }))))));
  }

}

ReactDOM.render(React.createElement(App, null), document.getElementById('app'));

},{}]},{},[1]);
