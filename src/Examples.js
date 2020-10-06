import React from 'react';


function Examples(props) {
    return (
        <table>
        <caption style={{marginBottom: "10px"}}>
          Примеры
        </caption>
        <thead>
          <th>Входные данные</th>
          <th>Выходные данные</th>
        </thead>
        <tbody>
          <tr>
            <td>3 2</td>
            <td>23<br /></td>
          </tr>
          <tr>
            <td>2 4</td>
            <td>-8<br /></td>
          </tr>
        </tbody>
      </table>
    )
}

export default Examples;