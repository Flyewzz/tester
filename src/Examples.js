import React from 'react';


function Examples(props) {

    const getSamples = (samples) => {
      samples = samples.split('|');
      let tags = [];
      for (let i = 0; i < samples.length; i += 2) {
        tags.push(
          <tr>
          <td>{ samples[i] }</td>
          <td>{ samples[i+1] }<br /></td>
        </tr>);
      }
      return tags;
    }

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
         {props.samples && getSamples(props.samples)}
        </tbody>
      </table>
    )
}

export default Examples;