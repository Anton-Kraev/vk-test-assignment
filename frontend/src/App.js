import React, { useEffect, useState } from 'react';

function App() {
  const [containers, setContainers] = useState([]);

  useEffect(() => {
    const fetchContainers = async () => {
      const response = await fetch('http://127.0.0.1:8080/containers');
      const data = await response.json();
      setContainers(data.containers);
    };

    fetchContainers().catch(e => console.log(e));
    const interval = setInterval(fetchContainers, 5000);
    return () => clearInterval(interval);
  }, []);

  return (
      <div>
        <h1>Container Status</h1>
        <table border={1}>
          <thead>
          <tr>
            <th>IP Адрес</th>
            <th>Последняя попытка пинга</th>
            <th>Последняя успешная попытка пинга</th>
            <th>Время ответа (ms)</th>
          </tr>
          </thead>
          <tbody>
          {containers.map((container) => (
              <tr key={container.id}>
                <td>{container.ip}</td>
                <td>{container.last_ping_attempt || '-'}</td>
                <td>{container.last_successful_ping || '-'}</td>
                <td>{container.response_time_ms || '-'}</td>
              </tr>
          ))}
          </tbody>
        </table>
      </div>
  );
}

export default App;
