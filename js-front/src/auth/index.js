import Cookies from 'universal-cookie';
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom';


export function isLoggedIn() {
    let result = new Cookies().get('sessionId');
    console.log(result);
    return Boolean(result);
}

export function logout(history) {
    new Cookies().remove('sessionId');
    history.push('/login');
}

export function PrivateRoute({ component: Component, ...rest }) {
    return (
        <Route {...rest} render={(props) => {
            if (isLoggedIn()) {
                return <Component {...props} />
            } else {
                return <Redirect to={{ pathname: '/login', state: { from: props.location } }} />
            }
        }} />
    )
}