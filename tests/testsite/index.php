<?php
// Set content type
header("Content-Type: text/html");

// Get the query parameter 'reflec'
$reflec = isset($_GET['reflec']) ? htmlspecialchars($_GET['reflec']) : "Not provided";
?>
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Test Server</title>
    <script>
        // Fetch external IP using JavaScript
        async function fetchExternalIP() {
            try {
                const response = await fetch('https://ifconfig.me/ip');
                const ip = await response.text();
                document.getElementById('external-ip').textContent = ip.trim();
            } catch (error) {
                document.getElementById('external-ip').textContent = 'Error fetching IP';
                console.error('Error fetching IP:', error);
            }
        }

        // Call the function on page load
        window.onload = fetchExternalIP;
    </script>
</head>
<body>
    <h1>Test Server</h1>
    <p>Your external IP address: <span id="external-ip">Loading...</span></p>
    <p>Reflec parameter value: <?php echo $reflec; ?></p>
</body>
</html>
