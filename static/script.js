var loginForm = document.getElementById('login');
var registerForm = document.getElementById('register');
var recoverForm = document.getElementById('recover');

function login() {
    loginForm.style.left = '4px';
    registerForm.style.right = '-520px';
    recoverForm.style.top = '-520px';

    loginForm.style.zIndex = 2;
    registerForm.style.zIndex = 1;
    recoverForm.style.zIndex = 0;
}

function register() {
    loginForm.style.left = '-520px';
    registerForm.style.right = '4px';
    recoverForm.style.top = '-520px';

    loginForm.style.zIndex = 1;
    registerForm.style.zIndex = 2;
    recoverForm.style.zIndex = 0;
}

function recover() {
    loginForm.style.top = '-520px';
    recoverForm.style.top = '4px';

    loginForm.style.zIndex = 1;
    registerForm.style.zIndex = 0;
    recoverForm.style.zIndex = 2;
}

function recoverLogin() {
    loginForm.style.top = '4px';
    recoverForm.style.top = '-520px';

    loginForm.style.zIndex = 2;
    registerForm.style.zIndex = 0;
    recoverForm.style.zIndex = 1;
}
