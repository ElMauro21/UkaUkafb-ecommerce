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

document.body.addEventListener('htmx:afterSwap', function (evt) {
    if (evt.detail.target.id === 'flash') {
        const flash = document.getElementById('flash');
        const type = flash?.getAttribute('data-type');

        // Only reset form if flash message indicates success
        if (type === 'success') {
            const form = document.getElementById('admin-products');
            form.reset();

            const addButton = document.getElementById('add-button');
            if (addButton) addButton.style.display = 'inline-block';

            const hiddenId = form.querySelector('[name="product-id"]');
            if (hiddenId) hiddenId.remove();
        }
    }
});

function fillProductForm(select) {
    const option = select.options[select.selectedIndex];
    const addButton = document.getElementById('add-button');

    if (!option.value) {
        document.getElementById('admin-products').reset();
        if (addButton) addButton.style.display = 'inline-block';

        const hiddenId = document.querySelector('[name="product-id"]');
        if (hiddenId) hiddenId.remove();

        return;
    }

    // Fill form fields
    document.querySelector('[name="product-name"]').value = option.dataset.name;
    document.querySelector('[name="product-description"]').value =
        option.dataset.description;
    document.querySelector('[name="product-weight"]').value =
        option.dataset.weight;
    document.querySelector('[name="product-size"]').value = option.dataset.size;
    document.querySelector('[name="product-price"]').value =
        option.dataset.price;
    document.querySelector('[name="product-quantity"]').value =
        option.dataset.quantity;
    document.querySelector('[name="product-image"]').value =
        option.dataset.image;

    let hiddenInput = document.querySelector('[name="product-id"]');
    if (!hiddenInput) {
        hiddenInput = document.createElement('input');
        hiddenInput.type = 'hidden';
        hiddenInput.name = 'product-id';
        document.getElementById('admin-products').appendChild(hiddenInput);
    }
    hiddenInput.value = option.value;

    if (addButton) addButton.style.display = 'none';
}

document
    .getElementById('admin-products')
    .addEventListener('submit', function (e) {
        const form = e.target;
        const submitter = document.activeElement;

        const productId = form.querySelector('[name="product-id"]')?.value;

        if (submitter?.value === 'Eliminar' && !productId) {
            e.preventDefault();
            alert('No hay ningún producto seleccionado para eliminar.');
        }
    });
