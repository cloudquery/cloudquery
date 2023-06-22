function detectOS() {
  const userAgent = navigator.userAgent.toLowerCase();

  if (userAgent.indexOf('win') >= 0) {
    return 'windows';
  } else if (userAgent.indexOf('mac') >= 0) {
    return 'mac';
  } else if (userAgent.indexOf('linux') >= 0) {
    return 'linux';
  }

  return 'Unknown';
}

document.addEventListener('DOMContentLoaded', function () {
  const os = detectOS();
  if (os !== 'mac') {
    // find class with name brew-install-command and hide it
    const brewInstallCommand = document.getElementsByClassName('brew-install-command');
    for (let i = 0; i < brewInstallCommand.length; i++) {
      brewInstallCommand[i].style.display = 'none';
    }
  }
}, false);
