import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import rehypeHighlight from 'rehype-highlight';
import 'highlight.js/styles/github.css';

interface Props {
  content: string;
}

export default function MdRenderer({ content }: Props) {
  return (
    <div className="markdown-body" style={{ fontSize: 15, lineHeight: 1.8 }}>
      <ReactMarkdown remarkPlugins={[remarkGfm]} rehypePlugins={[rehypeHighlight]}>
        {content}
      </ReactMarkdown>
    </div>
  );
}
