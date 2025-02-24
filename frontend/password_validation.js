// validatePassword.js
function validatePassword(password) {
    // Регулярное выражение для проверки:
    // - хотя бы одна строчная латинская буква (a-z)
    // - хотя бы одна заглавная латинская буква (A-Z)
    // - хотя бы одна цифра (0-9)
    // - хотя бы один спецсимвол (например, !@#$%^&*)
    const passwordPattern = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])(?=.*\W).{8,}$/;

    if (!passwordPattern.test(password)) {
        return "Пароль должен содержать минимум одну заглавную букву, одну строчную букву, одну цифру и один спецсимвол, а также быть длиной не менее 6 символов.";
    }
    return null;
}
