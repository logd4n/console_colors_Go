RUS:
Подключив данный файл к проекту, Вы можете изменять стиль и цвет текста.
Для изменения стиля текста просто вызывайте функцию SetColor(...), куда передаете необходимые значения.
Чтобы изменить цвет текста - передайте в SetColor переменные, названия которых начинаются с "text_".
Чтобы изменить цвет фона текста - передайте в SetColor переменные, названия которых начинаются с "background_".
Чтобы изменить атрибуты текста - передайте в SetColor переменные, названия которых начинаются с "attributes_".

ENG:
By connecting this file to the project, you can change the style and color of the text.
To change the text style, simply call the SetColor(...) function, where you pass the necessary values.
To change the text color, pass variables whose names begin with "text_" to SetColor.
To change the background color of the text, pass variables whose names begin with "background_" to SetColor.
To change text attributes, pass variables whose names begin with "attributes_" to SetColor.

Пример/Example:
SetColor(text_Green) //Устанавливаем зеленый цвет (Set the color to green)
println("Colors test\n") //В консоли текст зеленый (In the console, the text is green)