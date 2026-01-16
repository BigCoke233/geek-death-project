// Animate Hugo-built TOC
// author: Eltrac

document.addEventListener('DOMContentLoaded', () => {
  const toc = document.getElementById('toc');
  if (!toc) return;

  const headings = document.querySelectorAll('h2, h3');
  const observer = new IntersectionObserver(entries => {
    entries.forEach(entry => {
      if (!entry.isIntersecting) return;

      const id = entry.target.id;
      const tocLink = toc.querySelector(`a[href="#${id}"]`);
      if (!tocLink) return;

      const listItem = tocLink.parentElement;
      const isH2 = entry.target.tagName === 'H2';

      // Clear all active states when a new H2 is encountered
      if (isH2) {
        toc.querySelectorAll('li.active').forEach(el => el.classList.remove('active'));
      }

      // Set current item as active
      listItem.classList.add('active');

      // If this is an H3, find and activate its parent H2
      if (!isH2) {
        const parentH2Item = listItem.closest('ul').previousElementSibling?.closest('li');
        if (parentH2Item) parentH2Item.classList.add('active');
      }
    });
  }, {
    rootMargin: '0px 0px -50% 0px',
    threshold: 0.5
  });

  headings.forEach(heading => observer.observe(heading));
});
