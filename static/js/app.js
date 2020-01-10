import styled from 'styled-components'
import { ThemeProvider } from 'styled-components';
import { GlobalStyles } from './global';
import { theme } from './theme';
import { Home } from '../../src/components'

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

ReactDOM.render(
  <App />,
  document.getElementById('app')
);