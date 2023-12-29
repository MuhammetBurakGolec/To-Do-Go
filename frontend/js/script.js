document.getElementById('login-form').addEventListener('submit', function(e) {
    e.preventDefault();

    var username = document.getElementById('username').value;
    var password = document.getElementById('password').value;

    // Kullanıcı verilerini JSON formatında hazırla
    var userData = {
        user: username,
        password: password
    };

    // Backend sunucusuna POST isteği gönder
    fetch('localhost:8080/api/login', {  // '/api/login' sunucu tarafında uygun endpoint olmalıdır
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(userData)
    })
    .then(response => {
        if (response.ok) {
            return response.json();  // JWT token'ı almak için JSON olarak parse et
        } else {
            throw new Error('Login failed');
        }
    })
    .then(data => {
        // JWT token'ı cookie'ye kaydet
        document.cookie = "token=" + data.token + "; path=/";

        // Kullanıcıyı To-Do sayfasına yönlendir
        window.location.href = '/todo.html';  // Yönlendirilecek sayfa
    })
    .catch(error => {
        console.error('Error:', error);
        alert('Invalid username or password');
    });
});
