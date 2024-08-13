(function() {
    // Store original functions
    const originalAlert = window.alert;
    const originalConfirm = window.confirm;
    const originalPrompt = window.prompt;

    // Variables to track popup activity
    let popupDetected = false;
    const detectionTimeout = 5000; // 5 seconds timeout for detection

    // Utility function to show a custom message in the console
    function logXSSDetection(messageType, message, defaultInput) {
        console.log(
            `%cXSS Detected! Type: ${messageType} - Message: ${message} ${defaultInput ? 'Default Input: ' + defaultInput : ''}`,
            'color: green; font-weight: bold;'
        );
        popupDetected = true; // Mark that a popup was detected
    }

    function logNoXSSDetection() {
        if (!popupDetected) {
            console.log(
                `%cNo XSS Detected.`,
                'color: red; font-weight: bold;'
            );
        }
    }

    // Override alert
    window.alert = function(message) {
        logXSSDetection('Alert', message);
        // Call the original function to retain default behavior
        originalAlert(message);
    };

    // Override confirm
    window.confirm = function(message) {
        logXSSDetection('Confirm', message);
        // Call the original function to retain default behavior
        return originalConfirm(message);
    };

    // Override prompt
    window.prompt = function(message, defaultInput) {
        logXSSDetection('Prompt', message, defaultInput);
        // Call the original function to retain default behavior
        return originalPrompt(message, defaultInput);
    };

    // Set up a timeout to check for popup detection
    setTimeout(() => {
        logNoXSSDetection();
    }, detectionTimeout);

    // Optionally, you could use additional logic to continuously check or reset the popupDetected flag if needed.
})();
