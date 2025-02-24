function is_manager() {
    return fetch('http://localhost:8080/check_role', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('jwt')}`
        },
    })
    .then(response => response.ok)
    .catch(error => {
        console.error('Error to get role', error);
        alert('Error to get role');
        return false;
    });
}