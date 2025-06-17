const loginForm = document.getElementById('login');
const registerForm = document.getElementById('register');
const recoverForm = document.getElementById('recover');
const subMenu = document.getElementById('subMenu');
const toggleButton = document.getElementById('toggleMenuBtn');
const changeForm = document.getElementById('passwords');
const profileForm = document.getElementById('profile');

function change() {
    changeForm.style.right = '4px';
    profileForm.style.left = '-520px';

    changeForm.style.zIndex = 2;
    profileForm.style.zIndex = 1;
}

function profile() {
    changeForm.style.right = '-520px';
    profileForm.style.left = '4px';

    changeForm.style.zIndex = 1;
    profileForm.style.zIndex = 2;
}

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

toggleButton.addEventListener('click', function (event) {
    event.stopPropagation();
    subMenu.classList.toggle('open-menu');
});

document.addEventListener('click', function (event) {
    if (
        !subMenu.contains(event.target) &&
        !toggleButton.contains(event.target)
    ) {
        subMenu.classList.remove('open-menu');
    }
});
