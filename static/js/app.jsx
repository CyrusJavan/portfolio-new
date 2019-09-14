class App extends React.Component {
  render() {
    return (<Home />);
  }
}

class Home extends React.Component {
  render() {
    return (
      <div id="jumbotron-title" className="container jumbotron">
        <div className="row">
          <h1>Cyrus Javan</h1>
        </div>
        <div className="row">
          <h3>Hello. I do software development for a living.</h3>
        </div>
        <div className="row">
          <h3>I currently work at BetterHelp, the world's largest online mental health counseling platform.</h3>
        </div>
        <div className="row">
          <h3>Would communicating with me be potentially fruitful for both parties? Yes? Well that's wonderful, feel free to reach out on <a href="https://www.linkedin.com/in/cyrusjavan" rel="noopener noreferrer" target="_none">LinkedIn</a>.</h3>
        </div>
      </div>
    );
  }
}

ReactDOM.render(
  <App />,
  document.getElementById('app')
);