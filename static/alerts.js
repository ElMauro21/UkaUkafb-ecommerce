document.addEventListener('DOMContentLoaded', function () {
    const flash = document.getElementById('flash');
    const message = flash.dataset.message;
    const type = flash.dataset.type;

    if (message) {
        Swal.fire({
            toast: true,
            position: 'top-end',
            icon: type || 'info', // success, error, warning, info
            title: message,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
        });
    }
});
