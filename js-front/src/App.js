import LoginPage from './LoginPage';
import NotePage from './NotePage';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';
import { PrivateRoute } from './auth';

function App() {
  return (
    <div className="App py-5 px-1 px-md-3">
      <Router>
        <Route path="/" exact>
          <Redirect to="/login" />
        </Route>
        <Route path="/login" component={LoginPage} />
        <PrivateRoute path="/note" component={NotePage} />
      </Router>
    </div>
  );
}

export default App;
