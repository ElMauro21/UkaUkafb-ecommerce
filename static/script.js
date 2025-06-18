const loginForm = document.getElementById('login');
const registerForm = document.getElementById('register');
const recoverForm = document.getElementById('recover');
const subMenu = document.getElementById('subMenu');
const toggleButton = document.getElementById('toggleMenuBtn');
const changeForm = document.getElementById('passwords');
const profileForm = document.getElementById('profile');

function change() {
    profileForm.style.display = 'none';
    changeForm.style.display = 'block';
}

function profile() {
    profileForm.style.display = 'block';
    changeForm.style.display = 'none';
}

function login() {
    loginForm.style.display = 'block';
    registerForm.style.display = 'none';
    recoverForm.style.display = 'none';
}

function register() {
    loginForm.style.display = 'none';
    registerForm.style.display = 'block';
    recoverForm.style.display = 'none';
}

function recover() {
    loginForm.style.display = 'none';
    recoverForm.style.display = 'block';
}

function recoverLogin() {
    loginForm.style.display = 'block';
    recoverForm.style.display = 'none';
}

if (toggleButton && subMenu) {
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
}
