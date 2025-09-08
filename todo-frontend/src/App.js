import { useState } from "react";

export default function App() {
  const [quote, setQuote] = useState(null);
  const [error, setError] = useState(null);

  const fetchQuote = async () => {
    try {
      const res = await fetch("http://localhost:8080/quote");
      if (!res.ok) throw new Error("Failed to fetch quote");
      const data = await res.json();
      setQuote(data);
      setError(null);
    } catch (err) {
      setError("Error fetching quote");
    }
  };

  return (
    <div style={{ padding: "2rem", textAlign: "center" }}>
      <h1>Random Quote Generator</h1>
      <button
        onClick={fetchQuote}
        style={{
          margin: "1rem",
          padding: "0.5rem 1rem",
          border: "1px solid black",
          borderRadius: "5px",
          cursor: "pointer",
        }}
      >
        Get Quote
      </button>

      {error && <p style={{ color: "red" }}>{error}</p>}

      {quote && (
        <div style={{ marginTop: "1rem" }}>
          <p style={{ fontSize: "1.2rem" }}>"{quote.text}"</p>
          <p>— {quote.author}</p>
        </div>
      )}
    </div>
  );
}
