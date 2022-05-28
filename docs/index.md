---
title: "File Format Badger: make sure your files are valid!"
h1: File Format Badger
---
<script src="/js/popper.min.js"></script>
<script src="/js/bootstrap.bundle.min.js"></script>
<script>
    const popoverTriggerList = document.querySelectorAll('[data-bs-toggle="popover"]')
    const popoverList = [...popoverTriggerList].map(popoverTriggerEl => new bootstrap.Popover(popoverTriggerEl))
    console.log("running");
</script>
Badger is a linter <a class="badge rounded-pill bg-dark link-light text-decoration-none" href="https://en.wikipedia.org/wiki/Lint_(software)">?</a> for file formats. Are your files:

 * in the correct format?
 * with the correct extension?
 * with the correct image dimensions?
 * properly stripped of revealing metadata?
 * not too big or too small?
 * have decent names?

See [commands](/commands/index.html) for the list of supported file formats.

See [files](/files.html) for how to specify the files to check.

See [flags](/flags.html) for global flags that apply to multiple file formats, as well as details for complicated flag types like ranges and ratios.
