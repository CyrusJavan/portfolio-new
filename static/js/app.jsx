import styled from 'styled-components'
import { ThemeProvider } from 'styled-components';
import { GlobalStyles } from './global';
import { theme } from './theme';

class App extends React.Component {
  render() {
    return (
      <ThemeProvider theme={theme}>
        <>
          <GlobalStyles />
          <Home />
        </>
      </ThemeProvider>
    );
  }
}

class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="row">
          <div className="col-10 text-center pt-5 pl-5">
            <div >

              <img className="headshot" src="/static/img/headshot.jpg" alt="cyrus javan headshot"></img>
            </div>
            <h1 className="my-3">Cyrus Javan</h1>
            <h2 className="my-3">Software Engineer</h2>
            <h2 className="my-3">Currently helping build the world's largest online counseling platform at <a href="https://www.betterhelp.com">BetterHelp</a></h2>
          </div>
          <ul className="navbar-nav mr-auto">
            <li className="nav-item active">
              <a className="nav-link social-icon" href="https://github.com/CyrusJavan">
                <i className="fa fa-github"></i>
              </a>
            </li>
            <li className="nav-item">
              <a className="nav-link social-icon" href="mailto:javan.cyrus@gmail.com">
                <i className="fa fa-envelope"></i>

              </a>
            </li>
            <li className="nav-item">
              <a className="nav-link social-icon" href="https://www.linkedin.com/in/cyrusjavan">
                <i className="fa fa-linkedin"></i>
              </a>
            </li>
          </ul>

        </div>
      </div>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);