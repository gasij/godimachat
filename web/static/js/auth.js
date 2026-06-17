// auth.js — логика входа и регистрации

function showError(message) {
    const el = document.getElementById('error');
    el.textContent = message;
    el.classList.remove('hidden');
}

async function login(email, password) {
    try {
        const res = await fetch('/api/login', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
        });

        const data = await res.json();

        if (!res.ok) {
            showError(data.error || 'Ошибка входа');
            return;
        }

        localStorage.setItem('token', data.token);
        window.location.href = '/chat';
    } catch {
        showError('Не удалось подключиться к серверу');
    }
}

async function register(username, email, password) {
    try {
        const res = await fetch('/api/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ username, email, password }),
        });

        const data = await res.json();

        if (!res.ok) {
            showError(data.error || 'Ошибка регистрации');
            return;
        }

        localStorage.setItem('token', data.token);
        window.location.href = '/chat';
    } catch {
        showError('Не удалось подключиться к серверу');
    }
}
