import Cookies from 'universal-cookie';


export function isLoggedIn() {
    let result = new Cookies().get('sessionId');
    console.log(result);
    return Boolean(result);
}

export function logout(history) {
    new Cookies().remove('sessionId');
    history.push('/login');
}