var loginForm = document.getElementById('login');
var registerForm = document.getElementById('register');

function login() {
    loginForm.style.left = '4px';
    registerForm.style.right = '-520px';
}

function register() {
    loginForm.style.left = '-520px';
    registerForm.style.right = '4px';
}
