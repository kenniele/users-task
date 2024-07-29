document.addEventListener('DOMContentLoaded', function () {
    const filterForm = document.getElementById('filterForm');

    filterForm.addEventListener('submit', function (event) {
        event.preventDefault()

        const formData = new FormData(filterForm);
        const data = {};
        formData.forEach((value, key) => {
            data[key] = value;
        });

        const queryString = new URLSearchParams(data).toString();

        fetch(`/filter?${queryString}`, {
            method: 'GET',
            headers: {
                'Accept': 'text/html'
            }
        })
            .then(response => response.text())
            .then(html => {
                document.querySelector('.col-md-8').innerHTML = html;
            })
            .catch(error => {
                console.error('Error:', error);
            });
    });
});
