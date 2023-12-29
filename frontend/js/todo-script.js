// Cookie'den JWT token'ı al
function getCookie(name) {
    let cookie = {};
    document.cookie.split(';').forEach(function(el) {
        let [k,v] = el.split('=');
        cookie[k.trim()] = v;
    })
    return cookie[name];
}

// Yapılacaklar listesini sunucudan al
function fetchTodos() {
    fetch('/api/todos', {
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + getCookie('token') // JWT token'ı header'a ekle
        }
    })
    .then(response => response.json())
    .then(data => {
        updateTodoList(data.todos); // Listeyi güncelle
    })
    .catch(error => console.error('Error:', error));
}

// HTML listesini güncelle
function updateTodoList(todos) {
    const list = document.getElementById('todo-list');
    list.innerHTML = ''; // Listeyi temizle
    todos.forEach(todo => {
        let li = document.createElement('li');
        li.textContent = todo.task;
        list.appendChild(li);
    });
}

// Yeni todo ekleme
document.getElementById('todo-form').addEventListener('submit', function(e) {
    e.preventDefault();

    var newTask = document.getElementById('new-task').value;
    if (newTask) {
        // Sunucuya yeni todo ekleme isteği gönder
        fetch('/api/todos', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + getCookie('token')
            },
            body: JSON.stringify({ task: newTask })
        })
        .then(response => {
            if (response.ok) {
                fetchTodos(); // Todo listesini yeniden yükle
            }
        })
        .catch(error => console.error('Error:', error));

        document.getElementById('new-task').value = ''; // Input'u temizle
    }
});

// Sayfa yüklendiğinde todo'ları yükle
document.addEventListener('DOMContentLoaded', fetchTodos);
