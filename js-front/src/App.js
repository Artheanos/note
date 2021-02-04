import LoginPage from './LoginPage';
import NotePage from './NotePage';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';

function App() {
  return (
    <div className="App py-5">
      <Router>
        <Route path="/" exact>
          <Redirect to="/login" />
        </Route>
        <Route path="/login" component={LoginPage} />
        <Route path="/note" component={NotePage} />
      </Router>
    </div>
  );
}

export default App;
