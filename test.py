import base64
import os

def test_image_conversion():
    # Путь к тестовому изображению
    original_image = "image.png"
    restored_image = "restored.png"
    
    # 1. Конвертируем изображение в base64
    with open(original_image, "rb") as img_file:
        base64_str = base64.b64encode(img_file.read()).decode('utf-8')
    
    print(base64_str)
    print(f"Длина base64 строки: {len(base64_str)} символов")
    
    # 2. Конвертируем обратно в изображение
    with open(restored_image, "wb") as img_file:
        img_file.write(base64.b64decode(base64_str))
    
    # Проверяем размеры файлов
    original_size = os.path.getsize(original_image)
    restored_size = os.path.getsize(restored_image)
    
    print(f"Размер оригинального изображения: {original_size} байт")
    print(f"Размер восстановленного изображения: {restored_size} байт")
    print("Файлы идентичны:", original_size == restored_size)

# Запускаем тест
test_image_conversion()