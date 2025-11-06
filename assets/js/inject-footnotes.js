document.addEventListener("DOMContentLoaded", () => {
  const refs = Array.from(document.querySelectorAll("sup[id^='fnref']"));
  // reverse traversal, to keep injected footnotes in order
  for (let i = refs.length - 1; i >= 0; i--) {
    const ref = refs[i];

    const noteNumber = i + 1;
    const noteId = ref.id.replace("fnref", "fn");
    const note = document.getElementById(noteId);

    if (note) {
      // construct note element
      const inline = document.createElement("div");
      inline.className = "text-sm text-secondary flex gap-2 ml-4";

      const footnoteNumber = document.createElement('span')
      footnoteNumber.textContent = noteNumber;
      const footnoteContent = document.createElement('div');
      footnoteContent.innerHTML = note.innerHTML;

      inline.appendChild(footnoteNumber);
      inline.appendChild(footnoteContent);
      inline.id = noteId;

      // insert adjacent to paragraph
      // if it's inside a table, then insert outside the table
      if (ref.closest('table')) {
        ref.closest('table').insertAdjacentElement("afterend", inline);
      } else {
        ref.parentNode.insertAdjacentElement("afterend", inline);
      }
    }
  }
  // remove original footnotes
  document.getElementsByClassName("footnotes")[0]?.remove()
});
