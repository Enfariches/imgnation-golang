document.addEventListener("DOMContentLoaded", function () {
    // Функция для получения изображения
    function fetchImage() {

        // Создаем объект URL для разбора
        const fetchUrl = new URL(window.location.href);

        // Извлекаем ключ из пути
        const pathSegments = fetchUrl.pathname.split('/'); // Разбиваем путь по '/'
        const imageKey = pathSegments.pop(); // Получаем последний элемент в массиве

        if (!imageKey) {
            console.error('No image key provided in the URL');
            return;
        }
        const url = `/api/img/${imageKey}`;

        fetch(url)
            .then(response => {
                if (!response.ok) {
                    throw new Error('Network response was not ok' + response.statusText);
                }
                return response.blob(); // Конвертируем ответ в Blob
            })
            .then(imageBlob => {
                // Создаем объект URL для изображения
                const imageObjectURL = URL.createObjectURL(imageBlob);
                // Находим элемент изображения на странице
                document.getElementById("imageDisplay").src = imageObjectURL;
            })
            .catch(error => {
                console.error('There has been a problem with your fetch operation:', error);
            });
    }

    // Вызываем функцию после загрузки страницы
    fetchImage();
});