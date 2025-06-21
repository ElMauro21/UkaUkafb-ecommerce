document.addEventListener('DOMContentLoaded', () => {
    // ===== Form Toggles =====
    const loginForm = document.getElementById('login');
    const registerForm = document.getElementById('register');
    const recoverForm = document.getElementById('recover');
    const subMenu = document.getElementById('subMenu');
    const toggleButton = document.getElementById('toggleMenuBtn');
    const changeForm = document.getElementById('passwords');
    const profileForm = document.getElementById('profile');

    window.change = function () {
        profileForm?.style && (profileForm.style.display = 'none');
        changeForm?.style && (changeForm.style.display = 'block');
    };

    window.profile = function () {
        profileForm?.style && (profileForm.style.display = 'block');
        changeForm?.style && (changeForm.style.display = 'none');
    };

    window.login = function () {
        loginForm?.style && (loginForm.style.display = 'block');
        registerForm?.style && (registerForm.style.display = 'none');
        recoverForm?.style && (recoverForm.style.display = 'none');
    };

    window.register = function () {
        loginForm?.style && (loginForm.style.display = 'none');
        registerForm?.style && (registerForm.style.display = 'block');
        recoverForm?.style && (recoverForm.style.display = 'none');
    };

    window.recover = function () {
        loginForm?.style && (loginForm.style.display = 'none');
        recoverForm?.style && (recoverForm.style.display = 'block');
    };

    window.recoverLogin = function () {
        loginForm?.style && (loginForm.style.display = 'block');
        recoverForm?.style && (recoverForm.style.display = 'none');
    };

    if (toggleButton && subMenu) {
        toggleButton.addEventListener('click', event => {
            event.stopPropagation();
            subMenu.classList.toggle('open-menu');
        });

        document.addEventListener('click', event => {
            if (
                !subMenu.contains(event.target) &&
                !toggleButton.contains(event.target)
            ) {
                subMenu.classList.remove('open-menu');
            }
        });
    }

    setupAdminProductForm();
});

// ========== HTMX afterSwap Listener ==========
document.body.addEventListener('htmx:afterSwap', evt => {
    if (evt.detail.target.id === 'flash') {
        const flash = document.getElementById('flash');
        const type = flash?.getAttribute('data-type');

        if (type === 'success') {
            const form = document.getElementById('admin-products');
            form?.reset();

            const addButton = document.getElementById('add-button');
            if (addButton) addButton.style.display = 'inline-block';

            const hiddenId = form?.querySelector('[name="product-id"]');
            if (hiddenId) hiddenId.remove();
        }
    }

    if (evt.detail.target.id === 'admin-products') {
        setupAdminProductForm(); // rebind if form is swapped in
    }
});

// ========== Product Form Handling ==========
function setupAdminProductForm() {
    const form = document.getElementById('admin-products');
    const addButton = document.getElementById('add-button');

    if (!form) return;

    form.addEventListener('submit', e => {
        const submitter = document.activeElement;
        const productId = form.querySelector('[name="product-id"]')?.value;
    });

    window.fillProductForm = function (select) {
        const option = select.options[select.selectedIndex];
        if (!option.value) {
            form.reset();
            if (addButton) addButton.style.display = 'inline-block';

            const hiddenId = form.querySelector('[name="product-id"]');
            if (hiddenId) hiddenId.remove();
            return;
        }

        // Fill form fields
        form.querySelector('[name="product-name"]').value = option.dataset.name;
        form.querySelector('[name="product-description"]').value =
            option.dataset.description;
        form.querySelector('[name="product-weight"]').value =
            option.dataset.weight;
        form.querySelector('[name="product-size"]').value = option.dataset.size;
        form.querySelector('[name="product-price"]').value =
            option.dataset.price;
        form.querySelector('[name="product-quantity"]').value =
            option.dataset.quantity;
        form.querySelector('[name="product-image"]').value =
            option.dataset.image;
        form.querySelector('[name="product-image-two"]').value =
            option.dataset.image2;

        let hiddenInput = form.querySelector('[name="product-id"]');
        if (!hiddenInput) {
            hiddenInput = document.createElement('input');
            hiddenInput.type = 'hidden';
            hiddenInput.name = 'product-id';
            form.appendChild(hiddenInput);
        }

        hiddenInput.value = option.value;

        if (addButton) addButton.style.display = 'none';
    };
}

// Price formatted
document.addEventListener('DOMContentLoaded', () => {
    document.querySelectorAll('.price').forEach(el => {
        const raw = parseFloat(el.dataset.price);
        if (!isNaN(raw)) {
            el.textContent = raw.toLocaleString('es-CO', {
                style: 'currency',
                currency: 'COP',
                minimumFractionDigits: 2,
            });
        } else {
            el.textContent = 'Precio no disponible';
        }
    });
});

// Modal product

document.addEventListener('DOMContentLoaded', () => {
    const modalWrapper = document.getElementById('modal-wrapper');
    const modalOverlay = document.querySelector('.modal-overlay');
    const closeButton = document.getElementById('close');
    const openButtons = document.querySelectorAll('.open-modal');

    openButtons.forEach(button => {
        button.addEventListener('click', () => {
            const name = button.dataset.name;
            const description = button.dataset.description;
            const quantity = button.dataset.quantity;
            const price = parseFloat(button.dataset.price).toLocaleString(
                'es-CO',
                {
                    style: 'currency',
                    currency: 'COP',
                    minimumFractionDigits: 2,
                }
            );
            const image = button.dataset.image;

            document.querySelector(
                '#modal-container .modal-name h1'
            ).textContent = name;
            document.querySelector(
                '#modal-container .modal-description p'
            ).textContent = description;
            document.querySelector(
                '#modal-container .modal-quantity p'
            ).textContent = quantity;
            document.querySelector(
                '#modal-container .modal-price p'
            ).textContent = price;
            document.querySelector('#modal-container .modal-image img').src =
                image;

            modalWrapper.style.display = 'flex';
        });
    });

    if (modalOverlay) {
        modalOverlay.addEventListener('click', e => {
            if (e.target === modalOverlay) {
                modalWrapper.style.display = 'none';
            }
        });
    }

    if (closeButton) {
        closeButton.addEventListener('click', () => {
            modalWrapper.style.display = 'none';
        });
    }
});
