// chat.js — логика чата (WebSocket + REST API)

const token = localStorage.getItem('token');
if (!token) {
    window.location.href = '/login';
}

let ws = null;
let currentUser = null;

// === Инициализация ===

async function init() {
    await loadCurrentUser();
    await loadMessages();
    connectWebSocket();
    setupEventListeners();
}

async function loadCurrentUser() {
    const res = await fetch('/api/me', {
        headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) {
        localStorage.removeItem('token');
        window.location.href = '/login';
        return;
    }

    currentUser = await res.json();
    document.getElementById('username').textContent = currentUser.username;
}

async function loadMessages() {
    const res = await fetch('/api/messages', {
        headers: { Authorization: `Bearer ${token}` },
    });

    if (!res.ok) return;

    const messages = await res.json();
    const container = document.getElementById('messages');
    container.innerHTML = '';

    messages.forEach((msg) => appendMessage(msg));
    scrollToBottom();
}

// === WebSocket ===

function connectWebSocket() {
    const protocol = location.protocol === 'https:' ? 'wss:' : 'ws:';
    ws = new WebSocket(`${protocol}//${location.host}/api/ws`);

    ws.onopen = () => {
        document.getElementById('status').textContent = '● онлайн';
        document.getElementById('status').className = 'status online';
    };

    ws.onclose = () => {
        document.getElementById('status').textContent = '● офлайн';
        document.getElementById('status').className = 'status offline';
        // Переподключение через 3 секунды
        setTimeout(connectWebSocket, 3000);
    };

    ws.onmessage = (event) => {
        const payload = JSON.parse(event.data);
        if (payload.type === 'message') {
            appendMessage(payload.data);
            scrollToBottom();
        }
    };
}

// === Отправка сообщений ===

function setupEventListeners() {
    document.getElementById('messageForm').addEventListener('submit', async (e) => {
        e.preventDefault();
        const input = document.getElementById('messageInput');
        const content = input.value.trim();
        if (!content) return;

        const res = await fetch('/api/messages', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                Authorization: `Bearer ${token}`,
            },
            body: JSON.stringify({ content }),
        });

        if (res.ok) {
            input.value = '';
        }
    });

    document.getElementById('logoutBtn').addEventListener('click', () => {
        localStorage.removeItem('token');
        if (ws) ws.close();
        window.location.href = '/login';
    });
}

// === UI ===

function appendMessage(msg) {
    const container = document.getElementById('messages');
    const isOwn = currentUser && msg.user_id === currentUser.id;

    const div = document.createElement('div');
    div.className = `message${isOwn ? ' own' : ''}`;

    const time = new Date(msg.created_at).toLocaleTimeString('ru-RU', {
        hour: '2-digit',
        minute: '2-digit',
    });

    div.innerHTML = `
        <div class="message-header">
            <span class="message-author">${escapeHtml(msg.username)}</span>
            <span class="message-time">${time}</span>
        </div>
        <div class="message-content">${escapeHtml(msg.content)}</div>
    `;

    container.appendChild(div);
}

function scrollToBottom() {
    const main = document.querySelector('.chat-main');
    main.scrollTop = main.scrollHeight;
}

function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

init();
